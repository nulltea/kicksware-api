function initAutoSearch(actionUrl) {
	$(window).keypress(function(event) {
		if (event.keyCode === 13) {
			event.preventDefault();
			if ($(event.target).is("#auto-search")) {
				submitSearch(actionUrl, event.target.value);
			}
			return false;
		}
	});
}

function submitSearch(actionUrl, prefix) {
	if (isLoading()) return;
	resetSearch();
	$(".loading-overlay").toggleClass("load");
	$.ajax({
		url: `${actionUrl}?prefix=${prefix}`,
		type: "GET",
		contentType: "application/json",
		dataType: "json",
		success: displayResults
	});
}

function displayResults(data) {
	$(".search-grid").empty();

	if (!data) {
		notFoundHandle();
	}

	let len = Object.keys(data).length;

	if (len === 0) {
		notFoundHandle();
	}

	let index = 0;
	for (let item of data) {
		addSearchResultCell(item, ++index === len);
	}
}

function addSearchResultCell(item, last) {
	let searchCell = document.createElement("DIV");
	searchCell.className = "search-cell";
	searchCell.addEventListener("click", select);

	let thumb = document.createElement("DIV");
	thumb.addEventListener("click", select);
	thumb.className = "thumb";
	let image = document.createElement("IMG");
	image.setAttribute("src", item.imagePath);
	thumb.append(image);

	let info = document.createElement("DIV");
	info.addEventListener("click", select);
	info.className = "info";
	let brandTitle = document.createElement("SPAN");
	brandTitle.className = "brand-title";
	brandTitle.textContent = item.brandName;
	let modelTitle = document.createElement("SPAN");
	modelTitle.className = "model-title";
	modelTitle.textContent = item.modelName;
	let skuCode = document.createElement("SPAN");
	skuCode.className = "sku-code";
	skuCode.textContent = item.manufactureSku;
	let submitButton = createButton(item);
	info.append(brandTitle, modelTitle, skuCode, submitButton);

	searchCell.append(thumb, info);
	$(".search-grid").append(searchCell);

	if (last) {
		$(image).on("load", function (e) {
			e.stopPropagation();
			e.stopImmediatePropagation();
			$(".loading-overlay").toggleClass("load");
		});
	}
}

function resetSearch() {
	$(".search-grid.no-results").removeClass("no-results");
	$(".loading-overlay").css('display', 'block');
	$(".loading-overlay.load").removeClass("load");
	$(".no-result-content.active").removeClass("active");
	$(".selected").removeClass("selected");
	$(".form-actions button.next.active").removeClass("active");
	$(".form-actions button.next span").text("DETAILS");
}

function isLoading() {
	return $(".loading-overlay").hasClass("load");
}

function notFoundHandle() {
	$(".search-grid").toggleClass("no-results");
	$(".loading-overlay").hide();
	$(".loading-overlay.load").removeClass("load");
	$(".no-result-content").toggleClass("active");
	$(".form-actions button.next").toggleClass("active");
	$(".form-actions button.next span").text("START");
}

function select(item) {
	$(".selected").toggleClass("selected");
	$(item.target).closest(".search-cell").toggleClass("selected");
}

function createButton(model) {
	let submitButton = document.createElement("BUTTON");
	submitButton.type = "submit";
	submitButton.className = "button";

	//custom submit action handler
	$(submitButton).click(function (event) {
		window.location.href = `SetDetails?referenceID=${model.uniqueID}`;
		event.preventDefault();
	});

	let span = document.createElement("SPAN");
	span.textContent = "NEXT";

	let icon = document.createElementNS("http://www.w3.org/2000/svg", "svg");
	icon.setAttribute("xmlns", "http://www.w3.org/2000/svg");
	icon.setAttribute("viewBox", "0 0 24 24");
	let path = document.createElementNS("http://www.w3.org/2000/svg", "path");
	path.setAttribute("d", "M16 8v-4l8 8-8 8v-4h-16l8-8h8z");
	icon.append(path);

	submitButton.append(span, icon);
	return submitButton;
}

$(document).ready(function () {
	"use strict";
});
