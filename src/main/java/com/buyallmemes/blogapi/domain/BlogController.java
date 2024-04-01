package com.buyallmemes.blogapi.domain;

import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/posts")
@RequiredArgsConstructor
class BlogController {

    private final BlogPostsRetriever blogPostsRetriever;

    @GetMapping
    List<Post> getPosts() {
        return blogPostsRetriever.getAllPosts();
    }
}
