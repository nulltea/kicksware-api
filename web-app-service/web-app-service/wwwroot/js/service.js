function autocompleteInit(inp, arr) {

	var currentFocus;

	inp.addEventListener("input",
		function(e) {
			var b, i, val = this.value;

			closeAllLists();
			if (!val) {
				return false;
			}
			currentFocus = -1;

			var a = document.createElement("DIV");
			a.setAttribute("id", this.id + "autocomplete-list");
			a.setAttribute("class", "autocomplete-items");

			this.parentNode.appendChild(a);

			for (i = 0; i < arr.length; i++) {
	
				if (arr[i].substr(0, val.length).toUpperCase() == val.toUpperCase()) {
					b = document.createElement("DIV");

					b.innerHTML = "<strong>" + arr[i].substr(0, val.length) + "</strong>";
					b.innerHTML += arr[i].substr(val.length);
					b.innerHTML += "<input type='hidden' value='" + arr[i] + "'>";

					b.addEventListener("click",
						function() {
							inp.value = this.getElementsByTagName("input")[0].value;
							closeAllLists();
						});
					a.appendChild(b);
				}
			}
		});

	inp.addEventListener("keydown",
		function(e) {
			var x = document.getElementById(this.id + "autocomplete-list");
			if (x) x = x.getElementsByTagName("div");
			if (e.keyCode == 40) {
				currentFocus++;
				addActive(x);
			} else if (e.keyCode == 38) {
				currentFocus--;
				addActive(x);
			} else if (e.keyCode == 13) {
				e.preventDefault();
				if (currentFocus > -1) {
					if (x) x[currentFocus].click();
				}
			}
		});

	function addActive(x) {
		if (!x) return false;
		removeActive(x);
		if (currentFocus >= x.length) currentFocus = 0;
		if (currentFocus < 0) currentFocus = (x.length - 1);
		x[currentFocus].classList.add("autocomplete-active");
	}

	function removeActive(x) {

		for (var i = 0; i < x.length; i++) {
			x[i].classList.remove("autocomplete-active");
		}
	}

	function closeAllLists(element) {

		var x = document.getElementsByClassName("autocomplete-items");
		for (var i = 0; i < x.length; i++) {
			if (element != x[i] && element != inp) {
				x[i].parentNode.removeChild(x[i]);
			}
		}
	}
	document.addEventListener("click",
		function(e) {
			closeAllLists(e.target);
		});
}

function initCustomDropDown() {
	var customSelectors = document.getElementsByClassName("custom-select");
	for (var i = 0; i < customSelectors.length; i++) {
		var selectElements = customSelectors[i].getElementsByTagName("select")[0];
		/* For each element, create a new DIV that will act as the selected item: */
		var selectedItem = document.createElement("div");
		selectedItem.setAttribute("class", "select-selected");
		selectedItem.innerHTML = "";
		customSelectors[i].appendChild(selectedItem);
		/* For each element, create a new DIV that will contain the option list: */
		var itemBox = document.createElement("div");
		itemBox.setAttribute("class", "select-items select-hide");
		for (var j = 1; j < selectElements.length; j++) {
			/* For each option in the original select element,
			create a new DIV that will act as an option item: */
			var c = document.createElement("div");
			c.innerHTML = selectElements.options[j].innerHTML;
			c.addEventListener("click",
				function(e) {
					/* When an item is clicked, update the original select box,
					and the selected item: */
					var y, i, k;
					var originalSelect = this.parentNode.parentNode.getElementsByTagName("select")[0];
					var h = this.parentNode.previousSibling;
					for (i = 0; i < originalSelect.length; i++) {
						if (originalSelect.options[i].innerHTML == this.innerHTML) {
							originalSelect.selectedIndex = i;
							h.innerHTML = this.innerHTML;
							y = this.parentNode.getElementsByClassName("same-as-selected");
							for (k = 0; k < y.length; k++) {
								y[k].removeAttribute("class");
							}
							this.setAttribute("class", "same-as-selected");
							break;
						}
					}
					h.click();
				});
			itemBox.appendChild(c);
		}
		customSelectors[i].appendChild(itemBox);
		selectedItem.addEventListener("click",
			function(e) {
				/* When the select box is clicked, close any other select boxes,
				and open/close the current select box: */
				e.stopPropagation();
				closeAllSelect(this);
				this.nextSibling.classList.toggle("select-hide");
				this.classList.toggle("select-arrow-active");
			});
		selectedItem.innerHTML = $("select option:selected", customSelectors[i]).text();
	}

	function closeAllSelect(elmnt) {
		/* A function that will close all select boxes in the document,
		except the current select box: */
		var x, y, i, arrNo = [];
		x = document.getElementsByClassName("select-items");
		y = document.getElementsByClassName("select-selected");
		for (i = 0; i < y.length; i++) {
			if (elmnt == y[i]) {
				arrNo.push(i)
			} else {
				y[i].classList.remove("select-arrow-active");
			}
		}
		for (i = 0; i < x.length; i++) {
			if (arrNo.indexOf(i)) {
				x[i].classList.add("select-hide");
			}
		}
	}

	/* If the user clicks anywhere outside the select box,
	then close all select boxes: */
	document.addEventListener("click", closeAllSelect);
}


