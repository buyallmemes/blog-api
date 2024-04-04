package com.buyallmemes.blogapi;

import com.buyallmemes.blogapi.domain.Post;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class BlogApiApplicationTests {

    @Autowired
    private TestRestTemplate restTemplate;

    @Test
    void shouldReturnBlogPost() {
        Post[] response = restTemplate.getForObject("/posts", Post[].class);
        List<Post> posts = List.of(response);
        assertTrue(posts.size() > 1);

        assertAll(
                () -> assertEquals("20240331-lets-build.md", posts.get(posts.size() - 2)
                                                                  .filename()),
                () -> assertEquals("20240329-hello-world.md", posts.getLast()
                                                                   .filename()));

    }

}
