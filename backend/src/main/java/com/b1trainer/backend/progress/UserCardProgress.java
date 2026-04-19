package com.b1trainer.backend.progress;

import com.b1trainer.backend.card.Card;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import lombok.NoArgsConstructor;

import java.time.LocalDate;

@Entity
@Table(name = "user_card_progress")
@Getter
@Setter
@NoArgsConstructor
public class UserCardProgress {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "card_id", nullable = false)
    @JsonIgnoreProperties({"hibernateLazyInitializer", "handler"})
    private Card card;

    @Column(name = "repetition")
    private int repetition = 0;

    @Column(name = "easiness")
    private double easiness = 2.5;

    @Column(name = "interval_days")
    private int intervalDays = 1;

    @Column(name = "next_review_date")
    private LocalDate nextReviewDate = LocalDate.now();
}
