package com.buyallmemes.blogapi.domain;

import com.buyallmemes.blogapi.domain.model.Blog;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/posts")
@RequiredArgsConstructor
class BlogController {

    private final BlogConstructor blogConstructor;

    @GetMapping
    Blog getPosts() {
        return blogConstructor.get();
    }
}
