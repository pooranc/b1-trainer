package com.b1trainer.backend.algorithm;

import org.springframework.stereotype.Component;

@Component
public class SM2Algorithm {

    public SM2Result calculate(int repetition, double easiness, int intervalDays, int rating) {

        if (rating < 3) {
            return new SM2Result(0, easiness, 1);
        }

        double newEasiness = easiness + (0.1 - (5 - rating) * (0.08 + (5 - rating) * 0.02));
        newEasiness = Math.max(1.3, newEasiness);

        int newInterval;
        if (repetition == 0) {
            newInterval = 1;
        } else if (repetition == 1) {
            newInterval = 6;
        } else {
            newInterval = (int) Math.round(intervalDays * newEasiness);
        }
        return new SM2Result(repetition + 1, newEasiness, newInterval);
    }
}
