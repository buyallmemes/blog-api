package com.buyallmemes.blogapi.github;

import com.buyallmemes.blogapi.domain.dependencies.PostQueryRepository;
import com.buyallmemes.blogapi.domain.model.Post;
import com.buyallmemes.blogapi.github.dependencies.GitHubMDtoHTMLRenderer;
import com.buyallmemes.blogapi.mdparser.ParsedMD;
import lombok.RequiredArgsConstructor;
import org.kohsuke.github.GHContent;
import org.kohsuke.github.GHRepository;
import org.kohsuke.github.GitHub;

import java.io.IOException;
import java.io.InputStream;
import java.nio.charset.StandardCharsets;
import java.util.List;

@RequiredArgsConstructor
class GitHubPostRepository implements PostQueryRepository {

    public static final String REPOSITORY_PATH = "buyallmemes/blog-api";
    public static final String PATH_TO_POSTS = "posts";

    private final GitHub gitHubClient;
    private final GitHubMDtoHTMLRenderer htmlRenderer;

    @Override
    public List<Post> getAllPosts() {
        try {
            GHRepository repository = gitHubClient.getRepository(REPOSITORY_PATH);
            return repository.getDirectoryContent(PATH_TO_POSTS)
                             .stream()
                             .filter(ghContent -> ghContent.getName()
                                                           .endsWith(".md"))
                             .map(ghContent -> fetchPostsFromRepo(ghContent, repository))
                             .toList();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private Post fetchPostsFromRepo(GHContent ghContent, GHRepository repository) {
        String pathToFile = ghContent.getPath();
        try {
            GHContent fileContent = repository.getFileContent(pathToFile);
            return buildPost(fileContent);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private Post buildPost(GHContent fileContent) {
        try (InputStream inputStream = fileContent.read()) {
            byte[] content = inputStream
                    .readAllBytes();
            String mdContent = new String(content, StandardCharsets.UTF_8);
            ParsedMD parsedMD = htmlRenderer.renderHtml(mdContent);
            return Post.builder()
                       .filename(fileContent.getName())
                       .content(parsedMD.html())
                       .date(parsedMD.date())
                       .title(parsedMD.title())
                       .anchor(parsedMD.anchor())
                       .build();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

}
