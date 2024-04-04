package com.buyallmemes.blogapi.domain;

import lombok.Builder;
import lombok.NonNull;

@Builder
public record Post(@NonNull String filename,
                   @NonNull String content,
                   @NonNull String date,
                   @NonNull String title,
                   @NonNull String anchor) {
}
