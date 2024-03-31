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

    private final static String POSTS_PATH = "blog/posts";

    public List<Post> getAllPosts() {
        File[] files = getBlogPostFiles();
        return Arrays.stream(files)
                     .map(File::toPath)
                     .map(this::buildPost)
                     .sorted(this::reverseFileName)
                     .toList();
    }

    private File[] getBlogPostFiles() {
        URI uri = buildURI();
        return Paths.get(uri)
                    .toFile()
                    .listFiles();
    }

    private int reverseFileName(Post p1, Post p2) {
        return p2.filename()
                 .compareTo(p1.filename());
    }

    private URI buildURI() {
        try {
            return getClass().getClassLoader()
                             .getResource(POSTS_PATH)
                             .toURI();
        } catch (URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }

    private Post buildPost(Path path) {
        try {
            String fileName = path.getFileName()
                                  .toString();
            String content = Files.readString(path);
            return Post.builder()
                       .filename(fileName)
                       .content(content)
                       .build();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
