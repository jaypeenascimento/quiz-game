
# Prerequistes:
- [ ] Read the CSV
- [ ] Let user know it's right and wrong questions count.
- [ ] Iterate over all questions of a file.

# Additional info:
- File should default to problems.csv
- Question count will be small, so, no problem to keep it all at memory.
- Questions with invalid user responses should be considered wrong.
- Answers should be a single word/number.

## Como fazer:
- Ler o csv
- Printar a questao, numero da questao e quantidade de certas e erradas.
- Ler o input do usuario
- Validar Input x Resposta correta
- Incrementar count correto/incorreto
- Repetir ate finalizar.
- Printar resultado


## Estrutura:
- Struct quiz
  - Funcoes:
    - Printar Enunciado (x)
    - Obter resposta do usuario 
    - Validar resposta (x, res)
    - Incrementar Metricas (OK | MISS)
    - Printar resultado
  - Data
    - Questoes (Enunciado e respostas)
    - Metricas
