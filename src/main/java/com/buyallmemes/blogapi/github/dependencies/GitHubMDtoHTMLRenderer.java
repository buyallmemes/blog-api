package com.buyallmemes.blogapi.github.dependencies;

import com.buyallmemes.blogapi.mdparser.ParsedMD;

public interface GitHubMDtoHTMLRenderer {
    ParsedMD renderHtml(String mdContent);
}
