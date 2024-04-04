package com.buyallmemes.blogapi.mdparser;

import lombok.Builder;

@Builder
public record ParsedMD(String html, String title, String date, String anchor) {
}
