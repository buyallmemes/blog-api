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
        FrontMatter frontMatter = extractFrontMatter(parsed);
        String html = renderer.render(parsed);

        String title = frontMatter.title();
        String anchor = buildAnchor(title);
        String publishingDate = frontMatter.date();
        return ParsedMD.builder()
                       .html(html)
                       .title(title)
                       .date(publishingDate)
                       .anchor(anchor)
                       .build();
    }

    private FrontMatter extractFrontMatter(Node node) {
        YamlFrontMatterVisitor visitor = new YamlFrontMatterVisitor();
        node.accept(visitor);
        Map<String, List<String>> frontMatter = visitor.getData();
        String title = extractSimpleAttribute(frontMatter, "title");
        String date = extractSimpleAttribute(frontMatter, "date");
        return new FrontMatter(title, date);
    }

    private String extractSimpleAttribute(Map<String, List<String>> frontMatter, String attributeName) {
        return frontMatter.getOrDefault(attributeName, List.of(""))
                          .getFirst();
    }

    private record FrontMatter(String title, String date) {
    }

    private String buildAnchor(String title) {
        return title.toLowerCase()
                    .replaceAll("[^a-z0-9\\s-]", "")
                    .replaceAll("[\\s-]+", " ")
                    .trim()
                    .replaceAll("\\s+", "-");
    }
}
