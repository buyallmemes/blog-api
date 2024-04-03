package com.buyallmemes.blogapi.github;

import com.buyallmemes.blogapi.domain.dependencies.PostQueryRepository;
import com.buyallmemes.blogapi.github.dependencies.GitHubMDtoHTMLRenderer;
import org.kohsuke.github.GitHub;
import org.kohsuke.github.GitHubBuilder;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;
import java.util.Optional;

@Configuration
@ConditionalOnProperty(name = "blog.posts.fetch-from-github", havingValue = "true")
public class GitHubConfiguration {

    @Bean
    GitHub gitHub(@Value("${github.token}") Optional<String> token) throws IOException {
        GitHubBuilder gitHubBuilder = new GitHubBuilder();
        token.ifPresent(gitHubBuilder::withOAuthToken);
        return gitHubBuilder.build();
    }

    @Bean
    PostQueryRepository gitHubPostRepository(GitHub gitHub, GitHubMDtoHTMLRenderer htmlRenderer) {
        return new GitHubPostRepository(gitHub, htmlRenderer);
    }
}
