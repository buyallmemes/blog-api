package com.buyallmemes.blogapi.domain.dependencies;

import com.buyallmemes.blogapi.domain.model.Post;

import java.util.List;

public interface PostQueryRepository {
    List<Post> getAllPosts();
}
