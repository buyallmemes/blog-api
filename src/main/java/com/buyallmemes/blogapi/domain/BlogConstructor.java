package com.buyallmemes.blogapi.domain;

import com.buyallmemes.blogapi.domain.dependencies.PostQueryRepository;
import com.buyallmemes.blogapi.domain.model.Blog;
import com.buyallmemes.blogapi.domain.model.Post;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.function.Supplier;

@Component
@RequiredArgsConstructor
class BlogConstructor implements Supplier<Blog> {
    private final PostQueryRepository postQueryRepository;

    public Blog get() {
        List<Post> posts = fetchPosts();
        return new Blog(posts);
    }

    private List<Post> fetchPosts() {
        return postQueryRepository.getAllPosts()
                                  .stream()
                                  .sorted(this::reverseFileName)
                                  .toList();
    }

    private int reverseFileName(Post p1, Post p2) {
        return p2.filename()
                 .compareTo(p1.filename());
    }
}
