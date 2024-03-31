package com.buyallmemes.blogapi.domain;

import lombok.Builder;

@Builder
public record Post(String filename, String content) {
}
