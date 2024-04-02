package com.buyallmemes.blogapi;

import com.buyallmemes.blogapi.domain.Post;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class BlogApiApplicationTests {

    @Autowired
    private TestRestTemplate restTemplate;

    @Test
    void shouldReturnBlogPost() {
        Post[] response = restTemplate.getForObject("/posts", Post[].class);
        List<Post> posts = List.of(response);
        assertTrue(posts.size() > 1);

        assertEquals("31032024.md", posts.get(posts.size() - 2)
                                         .filename());
        assertEquals("29032024_hello_world.md", posts.getLast()
                                                     .filename());
    }

}
