package com.buyallmemes.blogapi;

import com.buyallmemes.blogapi.domain.Post;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.boot.test.web.server.LocalServerPort;

import static org.junit.jupiter.api.Assertions.assertNotNull;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class BlogApiApplicationTests {

    @Autowired
    private TestRestTemplate restTemplate;

    @Test
    void shouldReturnBlogPost() {
        Post[] posts = restTemplate.getForObject("/posts", Post[].class);
        assertNotNull(posts[0]);
    }

}
