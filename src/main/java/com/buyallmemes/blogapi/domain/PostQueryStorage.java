package com.buyallmemes.blogapi.domain;

import org.springframework.stereotype.Component;

import java.io.File;
import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Arrays;
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
        URI uri = buildURI();
        File[] files = Paths.get(uri)
                            .toFile()
                            .listFiles();
        return Arrays.stream(files)
                     .map(File::toPath)
                     .map(this::readString)
                     .toList();
    }

    private URI buildURI() {
        try {
            return getClass().getClassLoader()
                             .getResource("blog/posts")
                             .toURI();
        } catch (URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }

    private String readString(Path path) {
        try {
            return Files.readString(path);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
