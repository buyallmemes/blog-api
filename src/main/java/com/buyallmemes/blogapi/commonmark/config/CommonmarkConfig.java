package com.buyallmemes.blogapi.commonmark.config;

import org.commonmark.Extension;
import org.commonmark.ext.autolink.AutolinkExtension;
import org.commonmark.ext.front.matter.YamlFrontMatterExtension;
import org.commonmark.ext.gfm.tables.TablesExtension;
import org.commonmark.parser.Parser;
import org.commonmark.renderer.html.HtmlRenderer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.List;

@Configuration
public class CommonmarkConfig {
    @Bean
    List<Extension> commonmarkExtensions() {
        return List.of(
                TablesExtension.create(),
                AutolinkExtension.create(),
                YamlFrontMatterExtension.create()
        );
    }

    @Bean
    Parser parser(List<Extension> commonmarkExtensions) {
        return Parser.builder()
                     .extensions(commonmarkExtensions)
                     .build();
    }

    @Bean
    HtmlRenderer renderer(List<Extension> commonmarkExtensions) {
        return HtmlRenderer.builder()
                           .extensions(commonmarkExtensions)
                           .build();
    }
}
