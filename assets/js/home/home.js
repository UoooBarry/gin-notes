function focusNewInput(element) {
    const newInput = element.closest('#navbar-list').querySelector('input');
    if (newInput) {
        newInput.focus();
    }
}
