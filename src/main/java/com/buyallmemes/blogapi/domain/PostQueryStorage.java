package com.buyallmemes.blogapi.domain;

import org.springframework.stereotype.Component;

import java.io.File;
import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;

@Component
class PostQueryStorage {

    public List<Post> getAllPosts() {
        return getPostFilesContent().stream()
                                    .map(this::buildPost)
                                    .toList();

    }

    private Post buildPost(String content) {
        return Post.builder()
                   .content(content)
                   .build();
    }

    private List<String> getPostFilesContent() {
        try {
            URI uri = getClass().getClassLoader()
                                .getResource("blog/posts")
                                .toURI();
            Path path = Paths.get(uri);
            File[] files = path.toFile()
                               .listFiles();

            List<String> result = new ArrayList<>();
            for (File file : files) {
                result.add(Files.readString(file.toPath()));
            }
            return result;
        } catch (URISyntaxException | IOException e) {
            throw new RuntimeException(e);
        }
    }
}
