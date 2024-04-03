package com.buyallmemes.blogapi.commonmark;

import com.buyallmemes.blogapi.github.dependencies.GitHubMDtoHTMLRenderer;
import com.buyallmemes.blogapi.local.dependencies.LocalMDtoHTMLRenderer;
import lombok.RequiredArgsConstructor;
import org.commonmark.node.Node;
import org.commonmark.parser.Parser;
import org.commonmark.renderer.html.HtmlRenderer;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
class CommonmarkRenderer implements LocalMDtoHTMLRenderer, GitHubMDtoHTMLRenderer {

    private final Parser parser;
    private final HtmlRenderer renderer;

    @Override
    public String renderHtml(String mdContent) {
        Node parsed = parser.parse(mdContent);
        return renderer.render(parsed);
    }
}
