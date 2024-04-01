package com.buyallmemes.blogapi.domain;

import com.buyallmemes.blogapi.domain.dependencies.PostQueryRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Component;

import java.util.List;

@Component
@RequiredArgsConstructor
class BlogPostsRetriever {
    private final PostQueryRepository postQueryRepository;

    public List<Post> getAllPosts() {
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
