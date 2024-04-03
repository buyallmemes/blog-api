package com.buyallmemes.blogapi.local;

import com.buyallmemes.blogapi.domain.Post;
import com.buyallmemes.blogapi.domain.dependencies.PostQueryRepository;
import com.buyallmemes.blogapi.local.dependencies.LocalMDtoHTMLRenderer;
import lombok.RequiredArgsConstructor;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.stereotype.Component;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.List;

/**
 * This class is responsible for fetching blog posts from the local file system.
 * Intended to be used when the property `blog.posts.fetch-from-github` is set to `false`.
 * Mainly used for development.
 */
@Component
@RequiredArgsConstructor
@ConditionalOnProperty(name = "blog.posts.fetch-from-github", havingValue = "false", matchIfMissing = true)
class LocalPostRepository implements PostQueryRepository {

    private final static String POSTS_PATH = "posts";

    private final LocalMDtoHTMLRenderer htmlRenderer;

    @Override
    public List<Post> getAllPosts() {
        File[] files = findBlogPostFiles();
        return Arrays.stream(files)
                     .map(File::toPath)
                     .map(this::buildPost)
                     .sorted(this::reverseFileName)
                     .toList();
    }

    private File[] findBlogPostFiles() {
        return Paths.get("./" + POSTS_PATH)
                    .normalize()
                    .toAbsolutePath()
                    .toFile()
                    .listFiles();
    }

    private int reverseFileName(Post p1, Post p2) {
        return p2.filename()
                 .compareTo(p1.filename());
    }

    private Post buildPost(Path path) {
        try {
            String fileName = path.getFileName()
                                  .toString();
            String mdContent = Files.readString(path);
            String htmlContent = htmlRenderer.renderHtml(mdContent);
            return Post.builder()
                       .filename(fileName)
                       .content(htmlContent)
                       .build();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
