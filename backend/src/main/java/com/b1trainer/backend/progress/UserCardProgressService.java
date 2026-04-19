package com.b1trainer.backend.progress;

import com.b1trainer.backend.algorithm.SM2Algorithm;
import com.b1trainer.backend.algorithm.SM2Result;
import com.b1trainer.backend.card.Card;
import com.b1trainer.backend.card.CardRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDate;
import java.util.List;

@Service
@RequiredArgsConstructor
public class UserCardProgressService {

    private final UserCardProgressRepository progressRepository;
    private final CardRepository cardRepository;
    private final SM2Algorithm sm2Algorithm;

    public List<UserCardProgress> getDueCards() {
        return progressRepository.findByNextReviewDateLessThanEqual(LocalDate.now());
    }

    @Transactional
    public UserCardProgress submitRating(Long cardId, int rating) {
        System.out.println("Rating received: " + rating);
        UserCardProgress progress = progressRepository.findByCardId(cardId)
                .orElseThrow(() -> new RuntimeException("Progress not found for card: " + cardId));

        SM2Result result = sm2Algorithm.calculate(
                progress.getRepetition(),
                progress.getEasiness(),
                progress.getIntervalDays(),
                rating
        );

        progress.setRepetition(result.repetition());
        progress.setEasiness(result.easiness());
        progress.setIntervalDays(result.intervalDays());
        progress.setNextReviewDate(LocalDate.now().plusDays(result.intervalDays()));

        return progressRepository.save(progress);
    }

    @Transactional
    public UserCardProgress initProgress(Long cardId) {
        Card card = cardRepository.findById(cardId)
                .orElseThrow(() -> new RuntimeException("Card not found: " + cardId));

        UserCardProgress progress = new UserCardProgress();
        progress.setCard(card);
        return progressRepository.save(progress);
    }

    @Transactional
    public void initMissingProgress() {
        List<Card> allCards = cardRepository.findAll();
        for (Card card : allCards) {
            boolean exists = progressRepository.findByCardId(card.getId()).isPresent();
            if (!exists) {
                UserCardProgress progress = new UserCardProgress();
                progress.setCard(card);
                progressRepository.save(progress);
            }
        }
    }

}
