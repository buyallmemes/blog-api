package com.buyallmemes.blogapi.config;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.web.servlet.error.ErrorMvcAutoConfiguration;
import org.springframework.context.annotation.Bean;
import org.zalando.problem.jackson.ProblemModule;
import org.zalando.problem.violations.ConstraintViolationProblemModule;

@EnableAutoConfiguration(exclude = ErrorMvcAutoConfiguration.class)
public class ProblemConfig {
    @Bean
    public ObjectMapper objectMapper() {
        return new ObjectMapper().findAndRegisterModules()
                                 .registerModules(
                                         new ProblemModule(),
                                         new ConstraintViolationProblemModule());
    }
}
