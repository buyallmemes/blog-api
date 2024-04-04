package com.buyallmemes.blogapi;

import com.buyallmemes.blogapi.local.dependencies.LocalMDtoHTMLRenderer;
import com.buyallmemes.blogapi.mdparser.ParsedMD;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.junit.jupiter.api.Assertions.assertEquals;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class MarkdownRenderingTest {

    @Autowired
    private LocalMDtoHTMLRenderer htmlRenderer;

    @Test
    void shouldExtractFrontMatter() {
        String markdown = "---\n" +
                "title: \"Let's build 🚀\"\n" +
                "date: \"2021-08-01\"\n" +
                "---\n" +
                "## Hello World";
        ParsedMD actual = htmlRenderer.renderHtml(markdown);
        assertEquals("Let's build 🚀", actual.title());
        assertEquals("2021-08-01", actual.date());
        assertEquals("lets-build", actual.anchor());
    }
}
