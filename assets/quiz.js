// Shared multiple-choice quiz widget for SWE Interview Prep lessons.
// Usage:
// <div class="quiz" data-answer="1">
//   <div class="quiz-q">Question text?</div>
//   <button class="quiz-opt">Option A</button>
//   <button class="quiz-opt">Option B</button>
//   <div class="quiz-feedback"></div>
// </div>
// data-answer is the zero-based index of the correct .quiz-opt.
(function () {
  function initQuiz(quiz) {
    var correctIndex = parseInt(quiz.getAttribute('data-answer'), 10);
    var opts = quiz.querySelectorAll('.quiz-opt');
    var feedback = quiz.querySelector('.quiz-feedback');
    opts.forEach(function (opt, i) {
      opt.addEventListener('click', function () {
        if (opt.disabled) return;
        opts.forEach(function (o) { o.disabled = true; });
        opts[correctIndex].classList.add('correct');
        if (i !== correctIndex) {
          opt.classList.add('incorrect');
          if (feedback) feedback.textContent = 'Not quite — the highlighted option is correct.';
        } else {
          if (feedback) feedback.textContent = 'Correct.';
        }
      });
    });
  }
  document.addEventListener('DOMContentLoaded', function () {
    document.querySelectorAll('.quiz[data-answer]').forEach(initQuiz);
  });
})();
