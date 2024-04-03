package com.buyallmemes.blogapi.commonmark.config;

import org.commonmark.node.Heading;
import org.commonmark.renderer.html.CoreHtmlNodeRenderer;
import org.commonmark.renderer.html.HtmlNodeRendererContext;
import org.commonmark.renderer.html.HtmlWriter;

import java.util.LinkedHashMap;
import java.util.Map;

class HeadingRenderer extends CoreHtmlNodeRenderer {

    private final HtmlWriter html;
    private final HtmlNodeRendererContext context;

    HeadingRenderer(HtmlNodeRendererContext context) {
        super(context);
        this.context = context;
        this.html = context.getWriter();
    }

    @Override
    public void visit(Heading heading) {
        String htag = "h" + heading.getLevel();
        html.line();
        Map<String, String> attributes = context.extendAttributes(heading, htag, Map.of());
        String id = attributes.get("id");
        String sanitizedId = getSanitizedId(id);
        attributes.put("id", sanitizedId);

        html.tag(htag, attributes);

        visitChildren(heading);
        appendAnchor(sanitizedId);

        html.tag('/' + htag);
    }

    private String getSanitizedId(String id) {
        return id.replaceAll("\\W", " ")
                 .trim()
                 .replaceAll("\\s+", "-");
    }

    private void appendAnchor(String sanitizedId) {
        Map<String, String> anchorAttrs = new LinkedHashMap<>();
        anchorAttrs.put("href", "#" + sanitizedId);
        anchorAttrs.put("class", "anchor");
        html.tag("a", anchorAttrs);
        html.tag("/a");
    }
}
