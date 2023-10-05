package crigo

func generateRandomMouseMovements() RandomizeType {
	return RandomizeType(`var movement = document.querySelector('body');
    var movementposition = movement.getBoundingClientRect();
    var movementtotalx = movementposition.left + movementposition.right;
    var movementtotaly = movementposition.top + movementposition.bottom;
    var randox = Math.floor(Math.random() * (100 - 0 + 1)) + 0;
    var randoy = Math.floor(Math.random() * (100 - 0 + 1)) + 0;
    var movementx = movementtotalx / randox;
    var movementy = movementtotaly / randoy;
    var movementsecondevent = new MouseEvent('mousemove', {
      view: window,
      bubbles: true,
      cancelable: true,
      clientX: Math.round(movementx),
      clientY: Math.round(movementy),
    });
    movement.dispatchEvent(movementsecondevent);`)
}

func generateRandomScrollMovements() RandomizeType {
	return RandomizeType(`var rando = Math.floor(Math.random() * (10 - 0 + 1)) + 0;
    window.scroll({
      top:
        document.querySelector('body').getBoundingClientRect().bottom / rando,
      left: 0,
      behavior: 'smooth',
    });`)
}

func resetScrollMovement() RandomizeType {
	return RandomizeType(`var rando = Math.floor(Math.random() * (10 - 0 + 1)) + 0;
    window.scroll({
      top:
        document.querySelector('body').getBoundingClientRect().bottom / rando,
      left: 0,
      behavior: 'smooth',
    });`)
}
