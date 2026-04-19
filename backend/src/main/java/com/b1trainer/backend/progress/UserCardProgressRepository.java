package com.b1trainer.backend.progress;

import org.springframework.data.jpa.repository.JpaRepository;

import java.time.LocalDate;
import java.util.List;
import java.util.Optional;

public interface UserCardProgressRepository extends JpaRepository<UserCardProgress, Long> {

    List<UserCardProgress> findByNextReviewDateLessThanEqual(LocalDate date);

    Optional<UserCardProgress> findByCardId(Long cardId);
}