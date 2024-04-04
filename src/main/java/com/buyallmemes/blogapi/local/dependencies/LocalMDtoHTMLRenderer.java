package com.buyallmemes.blogapi.local.dependencies;

import com.buyallmemes.blogapi.mdparser.ParsedMD;

public interface LocalMDtoHTMLRenderer {
    ParsedMD renderHtml(String mdContent);
}
