;;<expr>              ::= <additive-expr> EOF
;;
;;<additive-expr>     ::= <multiplicative-expr> 
;;                     | <additive-expr> ('+' | '-') <multiplicative-expr>
;;
;;<multiplicative-expr> ::= <power-expr>
;;                       | <multiplicative-expr> ('*' | '/') <power-expr>
;;
;;<power-expr>        ::= <unary-expr>
;;                      | <unary-expr> '^' <power-expr>
;;
;;<unary-expr>        ::= '-' <unary-expr>
;;                      | <primary-expr>
;;
;;<primary-expr>      ::= <number>
;;                    | <identifier>
;;                      | '(' <additive-expr> ')'
;;
;;<number>            ::= <int>
;;                    | <float>
;;                      | <exp-number>
;;
;;<int>               ::= ['+' | '-']? <digit>+
;;
;;<float>             ::= ['+' | '-']? (<digit>+ '.' <digit>* 
;;                                    | '.' <digit>+)
;;
;;<exp-number>        ::= (<int> | <float>) ('e' | 'E') ['+' | '-']? <digit>+
;;
;;<identifier>        ::= <letter> <letter>*
;;
;;<digit>             ::= '0' | '1' | '2' | ... | '9'
;;
;;<letter>            ::= 'a' | 'b' | ... | 'z' 
;;                      | 'A' | 'B' | ... | 'Z'


(define (tokenize expr-string)
  (define (char-digit? c)
    (and (char? c) (>= (char->integer c) 48) (<= (char->integer c) 57)))
  (define (char-letter? c)
    (and (char? c)
         (or (and (>= (char->integer c) 65) (<= (char->integer c) 90))
             (and (>= (char->integer c) 97) (<= (char->integer c) 122)))))
  (define (char-whitespace? c)
    (and (char? c) (or (eq? c #\space) (eq? c #\tab) (eq? c #\newline))))
  (define (string->number-safe s)
    (and (> (string-length s) 0) (string->number s)))

  (define (sublist lst start end)
    (if (>= start end) '()
        (cons (list-ref lst start) (sublist lst (+ start 1) end))))

  (define (skip-number-chars chars pos)
    (if (>= pos (length chars)) pos
        (let ((c (list-ref chars pos)))
          (cond ((char-digit? c) (skip-number-chars chars (+ pos 1)))
                ((eq? c #\.) (skip-number-chars chars (+ pos 1)))
                ((or (eq? c #\e) (eq? c #\E))
                 (let ((pos1 (+ pos 1)))
                   (if (< pos1 (length chars))
                       (let ((c1 (list-ref chars pos1)))
                         (if (or (eq? c1 #\+) (eq? c1 #\-))
                             (skip-number-chars chars (+ pos1 1))
                             (skip-number-chars chars pos1)))
                       pos1)))
                (else pos)))))

  (define (skip-letter-chars chars pos)
    (if (>= pos (length chars)) pos
        (let ((c (list-ref chars pos)))
          (if (char-letter? c) (skip-letter-chars chars (+ pos 1)) pos))))

  (define (tokenize-helper chars pos tokens)
    (if (>= pos (length chars)) (reverse tokens)
        (let ((c (list-ref chars pos)))
          (cond
            ((char-whitespace? c) (tokenize-helper chars (+ pos 1) tokens))
            ((eq? c #\() (tokenize-helper chars (+ pos 1) (cons "(" tokens)))
            ((eq? c #\)) (tokenize-helper chars (+ pos 1) (cons ")" tokens)))
            ((eq? c #\+) (tokenize-helper chars (+ pos 1) (cons '+ tokens)))
            ((eq? c #\-) (tokenize-helper chars (+ pos 1) (cons '- tokens)))
            ((eq? c #\*) (tokenize-helper chars (+ pos 1) (cons '* tokens)))
            ((eq? c #\/) (tokenize-helper chars (+ pos 1) (cons '/ tokens)))
            ((eq? c #\^) (tokenize-helper chars (+ pos 1) (cons '^ tokens)))
            ((or (char-digit? c) (eq? c #\.))
             (let* ((start pos) (pos1 (skip-number-chars chars pos)))
               (if (eq? pos1 pos) #f
                   (let ((num-str (apply string (sublist chars start pos1))))
                     (let ((num (string->number-safe num-str)))
                       (if (and num (< pos1 (length chars))
                                   (char-letter? (list-ref chars pos1)))
                           #f
                           (if num (tokenize-helper chars pos1
                                                     (cons num tokens))
                               #f)))))))
            ((char-letter? c)
             (let* ((start pos) (pos1 (skip-letter-chars chars pos)))
               (if (< pos1 (length chars))
                   (if (char-digit? (list-ref chars pos1))
                       #f
                       (tokenize-helper chars pos1
                                        (cons (string->symbol
                                               (apply string
                                                      (sublist chars start pos1)))
                                              tokens)))
                   (tokenize-helper chars pos1
                                    (cons (string->symbol
                                           (apply string
                                                  (sublist chars start pos1)))
                                          tokens)))))
            (else #f)))))
  (let ((result (tokenize-helper (string->list expr-string) 0 '())))
    result))

(define (parse tokens)
  (define pos 0)
  (define token-list tokens)

  (define (current-token)
    (if (< pos (length token-list)) (list-ref token-list pos) #f))

  (define (advance) (set! pos (+ pos 1)))

  (define (match token)
    (if (equal? (current-token) token) (begin (advance) #t) #f))

  (define (is-value? tok)
    (and tok (or (number? tok)
                 (and (symbol? tok)
                      (not (member tok '(+ - * / ^)))))))

  (define (parse-expr)
    (let ((term (parse-term)))
      (if term (parse-expr-prime term) #f)))

  (define (parse-expr-prime left)
    (let ((op (current-token)))
      (if (member op '(+ -))
          (begin (advance)
                 (let ((right (parse-term)))
                   (if right (parse-expr-prime (list left op right)) #f)))
          left)))

  (define (parse-term)
    (let ((factor (parse-factor)))
      (if factor (parse-term-prime factor) #f)))

  (define (parse-term-prime left)
    (let ((op (current-token)))
      (if (member op '(* /))
          (begin (advance)
                 (let ((right (parse-factor)))
                   (if right (parse-term-prime (list left op right)) #f)))
          left)))

  (define (parse-factor)
    (let ((power (parse-power)))
      (if power (parse-factor-prime power) #f)))

  (define (parse-factor-prime left)
    (let ((op (current-token)))
      (if (eq? op '^)
          (begin (advance)
                 (let ((right (parse-power)))
                   (if right (list left op (parse-factor-prime right)) #f)))
          left)))

  (define (parse-power)
    (let ((tok (current-token)))
      (cond
        ((is-value? tok) (advance) tok)
        ((equal? tok "(")
         (advance)
         (let ((expr (parse-expr)))
           (if (and expr (match ")")) expr #f)))
        ((eq? tok '-)
         (advance)
         (let ((operand (parse-power)))
           (if operand (list '- operand) #f)))
        (else #f))))
  (if (null? tokens) #f
      (let ((tree (parse-expr)))
        (if (and tree (= pos (length token-list))) tree #f))))
(define (tree->scheme tree)
  (cond
    ((or (number? tree) (symbol? tree)) tree)
    ((list? tree)
     (cond
       ((and (= (length tree) 2) (eq? (car tree) '-))
        (let ((operand (tree->scheme (cadr tree))))
          (if operand (list '- operand) #f)))
       ((and (= (length tree) 3) (eq? (cadr tree) '^))
        (let ((left (tree->scheme (car tree)))
              (right (tree->scheme (caddr tree))))
          (if (and left right) (list 'expt left right) #f)))
       ((and (= (length tree) 3) (member (cadr tree) '(+ - * /)))
        (let ((left (tree->scheme (car tree)))
              (op (cadr tree))
              (right (tree->scheme (caddr tree))))
          (if (and left right) (list op left right) #f)))
       (else #f)))
    (else #f)))
