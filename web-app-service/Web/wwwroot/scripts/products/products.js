function initFilterPanel() {
	$(".toggle-menu").click(function () {
		let filterMenu = $(".filter-control");
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
		$(".spacer").height($(".accordion").height());
	});
}

function autocompleteFilter(inputSelector, filterValues) {
	let currentFocus;
	let input = $(inputSelector);
	let autocompleteList = $(".brand-list");
	input.on("input",function () {
		closeAllLists();

		let value = this.value;
		if (!value) {
			return false;
		}
		currentFocus = -1;

		for (let i = 0; i < filterValues.length; i++) {
			let filteredValue = filterValues[i];
			if (filteredValue.substr(0, value.length).toUpperCase() === value.toUpperCase()) {
				let id = `check-${filteredValue.replace(" ", "_").toLowerCase()}`;

				if ($(`#${id}`).length) {
					continue;
				}

				let brandRow = document.createElement("DIV");
				brandRow.classList.add("brand-row", "temp-row");

				let checkbox = document.createElement("INPUT");
				checkbox.type = "checkbox";
				checkbox.id = id;
				checkbox.className = "regular_checkbox";

				let label = document.createElement("LABEL");
				label.setAttribute("for", checkbox.id);
				label.innerHTML = `<strong>${filteredValue.substr(0, value.length)}</strong>`;
				label.innerHTML += filterValues[i].substr(value.length);
				label.innerHTML += `<input type='hidden' value='${filteredValue}'>`;

				brandRow.append(checkbox, label);
				$(checkbox).change(function () {
					$(brandRow).toggleClass("temp-row");
					closeAllLists();
				});
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
	rangeMin.on("input", function () {
		handleCollisionMin();
		syncMinPrice(this);
	});

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
}

function initIsotope() {
	let sortingButtons = $('.product_sorting_btn');
	let sortNums = $('.num_sorting_btn');
	let products = $(".product_grid");
	if (products.length) {
		let grid = products.isotope({
			itemSelector: '.product',
			layoutMode: 'fitRows',
			fitRows:
				{
					gutter: 30,
				},
			getSortData:
				{
					price: function (itemElement) {
						let priceEle = $(itemElement).find('.product_price').text().replace('$', '');
						return parseFloat(priceEle);
					},
					name: '.product_name',
					stars: function (itemElement) {
						let starsElements = $(itemElement).find('.rating');
						return starsElements.attr("data-rating");
					},
				},
			animationOptions:
				{
					duration: 750,
					easing: 'linear',
					queue: false,
				},
		});

		// Sort based on the value from the sorting_type dropdown
		sortingButtons.each(function () {
			$(this).on('click',
				function () {
					let parent = $(this).parent().parent().find('.sorting_text');
					parent.text($(this).text());
					let option = $(this).attr('data-isotope-option');
					option = JSON.parse(option);
					grid.isotope(option);
				});
		});
	}
}

$(document).ready(function () {
	initFilterPanel();

	priceRangeInit();
});