package com.buyallmemes.blogapi;

import com.buyallmemes.blogapi.local.dependencies.LocalMDtoHTMLRenderer;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class TestMarkDownRendering {

    @Autowired
    private LocalMDtoHTMLRenderer htmlRenderer;

    @ParameterizedTest
    @MethodSource("provideMarkdownContent")
    void shouldRenderMarkdownContent(String markdown, String expectedHtml) {
        String actual = htmlRenderer.renderHtml(markdown);
        assertEquals(expectedHtml, actual);
    }

    static Stream<Arguments> provideMarkdownContent() {
        return Stream.of(
                Arguments.of("## Hello World", "<h2 id=\"hello-world\">Hello World<a href=\"#hello-world\" class=\"anchor\"></a></h2>"),
                Arguments.of("## Let's build", "<h2 id=\"lets-build\">Let's build<a href=\"#lets-build\" class=\"anchor\"></a></h2>"),
                Arguments.of("## Let's build 🚀", "<h2 id=\"lets-build\">Let's build 🚀<a href=\"#lets-build\" class=\"anchor\"></a></h2>"),
                Arguments.of("### Java + Spring = ❤️", "<h3 id=\"java-spring\">Java + Spring = ❤️<a href=\"#java-spring\" class=\"anchor\"></a></h3>")
        );
    }
}
