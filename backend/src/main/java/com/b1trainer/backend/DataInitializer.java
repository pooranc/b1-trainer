package com.b1trainer.backend;

import com.b1trainer.backend.progress.UserCardProgressService;
import lombok.RequiredArgsConstructor;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor
public class DataInitializer implements CommandLineRunner {

    private final UserCardProgressService progressService;

    @Override
    public void run(String... args) {
        progressService.initMissingProgress();
        System.out.println("Progress initialized for all cards.");
    }
}