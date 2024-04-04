package com.buyallmemes.blogapi.commonmark;

import com.buyallmemes.blogapi.github.dependencies.GitHubMDtoHTMLRenderer;
import com.buyallmemes.blogapi.local.dependencies.LocalMDtoHTMLRenderer;
import com.buyallmemes.blogapi.mdparser.ParsedMD;
import lombok.RequiredArgsConstructor;
import org.commonmark.ext.front.matter.YamlFrontMatterVisitor;
import org.commonmark.node.Node;
import org.commonmark.parser.Parser;
import org.commonmark.renderer.html.HtmlRenderer;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.Map;

@Component
@RequiredArgsConstructor
class CommonmarkRenderer implements LocalMDtoHTMLRenderer, GitHubMDtoHTMLRenderer {

    private final Parser parser;
    private final HtmlRenderer renderer;

    @Override
    public ParsedMD renderHtml(String mdContent) {
        Node parsed = parser.parse(mdContent);
        YamlFrontMatterVisitor visitor = new YamlFrontMatterVisitor();
        parsed.accept(visitor);
        Map<String, List<String>> frontMatter = visitor.getData();
        String title = frontMatter.getOrDefault("title", List.of(""))
                                  .getFirst();
        String date = frontMatter.getOrDefault("date", List.of(""))
                                 .getFirst();
        String html = renderer.render(parsed);
        return ParsedMD.builder()
                       .html(html)
                       .title(title)
                       .date(date)
                       .anchor(sanitizeTitle(title))
                       .build();
    }

    private String sanitizeTitle(String title) {
        return title.toLowerCase()
                    .replaceAll("[^a-z0-9\\s-]", "")
                    .replaceAll("[\\s-]+", " ")
                    .trim()
                    .replaceAll("[\\s]", "-");
    }
}
