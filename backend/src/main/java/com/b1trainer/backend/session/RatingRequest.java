package com.b1trainer.backend.session;

public record RatingRequest(
        Long cardId,
        int rating
) {
}
