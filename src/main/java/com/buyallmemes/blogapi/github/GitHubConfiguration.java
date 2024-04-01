package com.buyallmemes.blogapi.github;

import org.kohsuke.github.GitHub;
import org.kohsuke.github.GitHubBuilder;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.io.IOException;

@Configuration
public class GitHubConfiguration {

    @Bean
    public GitHub gitHub(@Value("${github.token}") String token) throws IOException {
        return new GitHubBuilder().withOAuthToken(token)
                                  .build();
    }

}
