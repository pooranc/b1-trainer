package com.b1trainer.backend.session;

import com.b1trainer.backend.progress.UserCardProgress;
import com.b1trainer.backend.progress.UserCardProgressService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/session")
@RequiredArgsConstructor
@CrossOrigin(origins = "*")
public class SessionController {

    private final UserCardProgressService progressService;

    @GetMapping("/due")
    public ResponseEntity<List<UserCardProgress>> getDueCards() {
        return ResponseEntity.ok(progressService.getDueCards());
    }

    @PostMapping("/rate")
    public ResponseEntity<UserCardProgress> submitRating(@RequestBody RatingRequest request) {
        UserCardProgress updated = progressService.submitRating(request.cardId(), request.rating());
        return ResponseEntity.ok(updated);
    }

    @PostMapping("/init/{cardId}")
    public ResponseEntity<UserCardProgress> initCard(@PathVariable Long cardId) {
        UserCardProgress progress = progressService.initProgress(cardId);
        return ResponseEntity.ok(progress);
    }
}
