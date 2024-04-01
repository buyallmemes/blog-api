package com.buyallmemes.blogapi.github;

import org.kohsuke.github.GitHub;
import org.kohsuke.github.GitHubBuilder;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;
import java.util.Optional;

@Configuration
public class GitHubConfiguration {

    @Bean
    public GitHub gitHub(@Value("${github.token}") Optional<String> token) throws IOException {
        GitHubBuilder gitHubBuilder = new GitHubBuilder();
        token.ifPresent(gitHubBuilder::withOAuthToken);
        return gitHubBuilder.build();
    }

}
