function autocompleteInit(input, dataValues) {
	let currentFocus;

	input.addEventListener("input", function() {
		let value = this.value;

		closeAllLists();
		if (!value) {
			return false;
		}
		currentFocus = -1;

		let autocompleteList = document.createElement("DIV");
		autocompleteList.setAttribute("id", this.id + "autocomplete-list");
		autocompleteList.setAttribute("class", "autocomplete-items");

		this.parentNode.appendChild(autocompleteList);

		for (let i = 0; i < dataValues.length; i++) {

			if (dataValues[i].substr(0, value.length).toUpperCase() === value.toUpperCase()) {
				let item = document.createElement("DIV");

				item.innerHTML = `<strong>${dataValues[i].substr(0, value.length)}</strong>`;
				item.innerHTML += dataValues[i].substr(value.length);
				item.innerHTML += `<input type='hidden' value='${dataValues[i]}'>`;

				item.addEventListener("click",
					function() {
						input.value = this.getElementsByTagName("input")[0].value;
						closeAllLists();
					});
				autocompleteList.appendChild(item);
			}
		}
	});

	input.addEventListener("keydown",
		function(e) {
			let selectedList = document.getElementById(this.id + "autocomplete-list");
			if (selectedList) selectedList = selectedList.getElementsByTagName("div");
			if (e.keyCode === 40) {
				currentFocus++;
				addActive(selectedList);
			} else if (e.keyCode === 38) {
				currentFocus--;
				addActive(selectedList);
			} else if (e.keyCode === 13) {
				e.preventDefault();
				if (currentFocus > -1) {
					if (selectedList) selectedList[currentFocus].click();
				}
			}
		});

	function addActive(item) {
		if (!item) return false;
		removeActive(item);
		if (currentFocus >= item.length) currentFocus = 0;
		if (currentFocus < 0) currentFocus = (item.length - 1);
		item[currentFocus].classList.add("autocomplete-active");
	}

	function removeActive(item) {
		for (let i = 0; i < item.length; i++) {
			item[i].classList.remove("autocomplete-active");
		}
	}

	function closeAllLists(element) {
		let x = document.getElementsByClassName("autocomplete-items");
		for (let i = 0; i < x.length; i++) {
			if (element !== x[i] && element !== input) {
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
	let customSelectors = document.getElementsByClassName("list-select");
	for (let i = 0; i < customSelectors.length; i++) {
		let selectElement = customSelectors[i].getElementsByTagName("select")[0];

		let selectedItem = document.createElement("div");
		selectedItem.setAttribute("class", "select-selected");
		selectedItem.innerHTML = "";
		customSelectors[i].appendChild(selectedItem);

		let itemBox = document.createElement("div");
		itemBox.setAttribute("class", "select-items select-hide");
		for (let j = 1; j < selectElement.length; j++) {
			let option = document.createElement("div");
			option.innerHTML = selectElement.options[j].innerHTML;
			option.addEventListener("click",
				function(e) {
					let originalSelect = this.parentNode.parentNode.getElementsByTagName("select")[0];
					let previousSibling = this.parentNode.previousSibling;
					for (let i = 0; i < originalSelect.length; i++) {
						if (originalSelect.options[i].innerHTML === this.innerHTML) {
							originalSelect.selectedIndex = i;
							previousSibling.innerHTML = this.innerHTML;
							let sameSelected = this.parentNode.getElementsByClassName("same-as-selected");
							for (let k = 0; k < sameSelected.length; k++) {
								sameSelected[k].removeAttribute("class");
							}
							this.setAttribute("class", "same-as-selected");
							break;
						}
					}
					previousSibling.click();
					$(selectElement).change();
				});
			itemBox.appendChild(option);
		}
		customSelectors[i].appendChild(itemBox);
		selectedItem.addEventListener("click",
			function(e) {
				e.stopPropagation();
				closeAllSelect(this);
				$(this.nextSibling).toggleClass("select-hide");
				$(this).toggleClass("select-arrow-active");
				$(this.parentElement).toggleClass("dropped");
			});
		selectedItem.innerHTML = $("select option:selected", customSelectors[i]).text();
	}

	function closeAllSelect(element) {
		let numArray  = [];
		let selectItems = document.getElementsByClassName("select-items");
		let selectedItem = document.getElementsByClassName("select-selected");
		for (let i = 0; i < selectedItem.length; i++) {
			if (element === selectedItem[i]) {
				numArray.push(i)
			} else {
				$(selectedItem[i]).removeClass("select-arrow-active");
				$(selectedItem[i].parentElement).removeClass("dropped");
			}
		}
		for (let i = 0; i < selectItems.length; i++) {
			if (numArray.indexOf(i)) {
				selectItems[i].classList.add("select-hide");
			}
		}
	}

	// Handle outside the select box click
	document.addEventListener("click", closeAllSelect);
}

function isDescendant(parent, child) {
	if (child.id === "login" || child.id === "sing-up") return true;

	let node = child.parentNode;
	while (node !== null) {
		if (node === parent) {
			return true;
		}
		node = node.parentNode;
	}
	return false;
}

function toggleLoadOverlay() {
	let preloader = $(".preloader")
	$('.modal-backdrop').toggleClass("show");
	if (preloader.is(":visible")) {
		preloader.modal("hide");
		preloader.hide();
		$('body').removeClass("modal-open");
	} else {
		preloader.show();
		preloader.modal("show");
	}
}


