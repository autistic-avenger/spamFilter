# Email Spam Classifier

A Naive Bayes spam classifier written in Go.

## How It Works

The classifier follows the **Naive Bayes** approach:

1. Parse emails from the training dataset.
2. Tokenize each email into words.
3. Count how frequently each word appears in Spam and Ham emails.
4. Calculate the probability of an email being Spam or Ham.
5. Classify the email based on the higher probability.


## Pipeline

```mermaid
graph TD
    A[Training Emails] --> B[Spam]
    A --> C[Ham]

    B --> D[Tokenization]
    C --> E[Tokenization]

    D --> F[Word Counts]
    E --> G[Word Counts]

    F --> H[Naive Bayes Model]
    G --> H

    H --> I[New Email]
    I --> J[Tokenization]

    J --> K[Spam Probability]
    J --> L[Ham Probability]

    K --> M[Higher Score]
    L --> M

    M --> N[Spam / Ham]
```

## Tech Stack

* **Language:** Go
* **Dataset:** [Enron Email Dataset](https://www.kaggle.com/datasets/wcukierski/enron-email-dataset)
* **Algorithm:** Naive Bayes



