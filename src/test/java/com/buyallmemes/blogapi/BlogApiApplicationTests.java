package com.buyallmemes.blogapi;

import com.buyallmemes.blogapi.domain.Post;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;

import static org.junit.jupiter.api.Assertions.assertEquals;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class BlogApiApplicationTests {

    @Autowired
    private TestRestTemplate restTemplate;

    @Test
    void shouldReturnBlogPost() {
        Post[] posts = restTemplate.getForObject("/posts", Post[].class);
        assertEquals(2, posts.length);

        assertEquals("31032024.md", posts[0].filename());
        assertEquals("29032024_hello_world.md", posts[1].filename());
    }

}
