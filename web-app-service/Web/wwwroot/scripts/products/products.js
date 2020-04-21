﻿function initFilterPanel() {
	let filterMenu = $(".filter-sidebar");
	$(".toggle-menu").click(function () {

		let accordion = $(".accordion");

		if (filterMenu.hasClass("active")) {
			accordion.toggleClass("collapsed");
			$(".toggle-menu span").text("SHOW FILTERS");
			$(this).toggleClass("pressed");
		} else {
			$(".toggle-menu span").text("HIDE FILTERS");
			$(this).toggleClass("pressed");
			accordion.toggleClass("collapsed");
		}

		filterMenu.toggleClass("active");
	});

	$(".accordion-control").change(function (e) {
		filterMenu.find(".spacer").height($(".accordion").height());
	});
}

function autocompleteFilter(inputSelector, filterValues) {
	let currentFocus;
	let input = $(inputSelector);
	input.on("input",function () {
		closeAllLists();

		let value = this.value;
		if (!value) {
			return false;
		}
		currentFocus = -1;

		for (let i = 0; i < filterValues.length; i++) {
			let filteredValue = filterValues[i];
			if (filteredValue["Caption"].substr(0, value.length).toUpperCase() === value.toUpperCase()) {
				let id = filterValues[i]["RenderId"];

				if ($(`#${id}`).length) {
					continue;
				}

				let brandRow = document.createElement("DIV");
				brandRow.classList.add("brand-row", "temp-row");

				let checkbox = document.createElement("INPUT");
				checkbox.type = "checkbox";
				checkbox.id = id;
				checkbox.value = filteredValue["Value"];
				checkbox.className = "regular_checkbox";

				let label = document.createElement("LABEL");
				label.setAttribute("for", id);
				label.innerHTML = `<strong>${filteredValue["Caption"].substr(0, value.length)}</strong>`;
				label.innerHTML += filteredValue["Caption"].substr(value.length);

				brandRow.append(checkbox, label);
				$(checkbox).change(function () {
					$(brandRow).toggleClass("temp-row");
					closeAllLists();
				});
				chipsInit($(checkbox));
				bindFilterInputSubmitEvent($(checkbox));
				$(".brand-list").prepend($(brandRow));
			}
		}
	});

	function closeAllLists() {
		$(".temp-row").remove();
		$(".brand-row strong").each(function(){
			$(this).replaceWith($(this).text());
		});
	}

	document.addEventListener("click",
	function (e) {
		if (!$(e.target).closest(".brand-row").length){
			closeAllLists();
		}
	});
}

function priceRangeInit() {
	let rangeMax = $(".price-slider.max");
	let rangeMin = $(".price-slider.min");
	let inputMax = $(".price-input.max");
	let inputMin = $(".price-input.min");
	let maxRangeElement = rangeMax.get(0);
	let minRangeElement = rangeMin.get(0);

	function syncMaxPrice(source) {
		if (source.type === "range") {
			inputMax.val(rangeMax.val());
		} else {
			maxRangeElement.val(inputMax.val());
		}
	}
	function syncMinPrice(source) {
		if (source.type === "range") {
			inputMin.val(rangeMin.val());
		} else {
			rangeMin.val(inputMin.val());
		}
	}

	function round(num) {
		return parseInt(Math.round(num * 0.2, 0) * 5);
	}

	let minValue = parseInt(minRangeElement.value);
	let maxValue = parseInt(maxRangeElement.value);
	let offset = (maxRangeElement.max - maxRangeElement.min) * 4 / 35;

	function handleCollisionMin() {
		minValue = parseInt(minRangeElement.value);
		maxValue = parseInt(maxRangeElement.value);
		offset = (maxRangeElement.max - maxRangeElement.min) * 4 / 35;


		if (minValue > maxValue - offset) {
			maxRangeElement.value = minValue + offset;
			syncMaxPrice(maxRangeElement);
			if (maxValue === parseInt(maxRangeElement.max)) {
				minRangeElement.value = round(parseInt(maxRangeElement.max) - offset);
			}
		}
	}
	function handleCollisionMax() {
		minValue = parseInt(minRangeElement.value);
		maxValue = parseInt(maxRangeElement.value);
		offset = (maxRangeElement.max - maxRangeElement.min) * 4 / 35;

		if (maxValue < minValue + offset) {
			minRangeElement.value = maxValue - offset;
			syncMinPrice(minRangeElement);
			if (minValue === parseInt(minRangeElement.min)) {
				maxRangeElement.value = round(parseInt(maxRangeElement.min) + offset);
			}
		}
	}

	rangeMax.on("input", function () {
		handleCollisionMax();
		syncMaxPrice(this);
	});
	bindFilterInputSubmitEvent(rangeMax, "mouseup");
	rangeMin.on("input", function () {
		handleCollisionMin();
		syncMinPrice(this);
	});
	bindFilterInputSubmitEvent(rangeMin, "mouseup");

	inputMax.on("input", function () {
		handleCollisionMax();
		syncMaxPrice(this);
	});
	inputMin.on("input", function () {
		handleCollisionMin();
		syncMinPrice(this);
	});

	syncMaxPrice(maxRangeElement);
	syncMinPrice(minRangeElement);
} //TODO bind submit action on range change

function chipsInit(option = null){
	let filterOverbar = $(".filter-overbar");
	let chipsPanel = $(".filter-chips");
	let filterOptions = option ?? $(".section-content input[type=checkbox]");
	filterOptions.each(function () {
		$(this).change(function () {
			showChipsPanel();
			let option = $(this);
			let id = `chip-${this.id}`;

			if (!option.is(":checked")) {
				$(`#${id}`).remove();
				hideChipsPanel();
				return;
			}

			let label = option.find("+ label");
			let chip = document.createElement("SPAN");
			chip.id = id;
			chip.className = "chip";
			chip.textContent = label.text();

			let close = document.createElement("BUTTON");
			close.className = "close-button";
			$(close).click(function() {
				$(chip).remove();
				option.prop('checked', false);
				hideChipsPanel();
			});
			bindFilterInputSubmitEvent($(close), "click");
			chip.append(close);
			chipsPanel.append(chip);
		});
	});

	function showChipsPanel() {
		if (!filterOverbar.hasClass("active")) {
			filterOverbar.toggleClass("active");
		}
	}

	function hideChipsPanel() {
		if (!$(".chip").length) {
			filterOverbar.toggleClass("active");
		}
	}

	function resetAllFilters() {
		filterOptions.each(function () {
			$(this).prop('checked', false);
		});
		$(".chip").remove();
		hideChipsPanel();
	}
	if(option == null){
		let resetButton = $("#filter-reset");
		resetButton.click(resetAllFilters);
		bindFilterInputSubmitEvent(resetButton, "click");
	}
}

function layoutToggleInit() {
	let toggleWrapper = $(".layout-toggle");
	let toggleInput = toggleWrapper.find("input[type=checkbox]");

	toggleInput.change(function () {
		let productsView = $(".products-view");
		productsView.toggleClass("changing");
		setTimeout(function () {
			productsView.toggleClass("grid").toggleClass("list");
			productsView.toggleClass("changing");
		}, 500);
	})
}

function filterNavigatorInit(){
	bindFilterInputSubmitEvent($(".filter-sidebar input[id^=filter-control]"));
}


function bindFilterInputSubmitEvent(element, event="change") {
	element.on(event, function () {
		$.post("/Shop/RequestFilter", {filterInputs: formFilterParameters(), sortBy: formSortParameter() }, function(response) {
			$(".result-content").html(response["content"]);
			$(".count span").text(`Showing 0-${Math.min(response["pageSize"], response["length"])} / ${response["length"]} results`);
		});
	});

	function formFilterParameters() {
		let formMap = {};
		$(".filter-sidebar input[id^=filter-control]").each(function () {
			let param = {};
			let checked = this.type === "checkbox" ? $(this).is(":checked") : true;
			let value = this.value;
			if (this.type === "number" || this.type === "range") {
				value = parseFloat(value)
			}
			param[checked] = JSON.stringify(value);
			formMap[this.id] = param;
		});
		return formMap;
	}

	function formSortParameter() {
		return $(".sort_type select").val();
	}
}

function sortingInit(){
	bindFilterInputSubmitEvent($(".sort_type select"));
}

$(document).ready(function () {
	initFilterPanel();

	priceRangeInit();

	chipsInit();

	layoutToggleInit();

	filterNavigatorInit();

	sortingInit();
});