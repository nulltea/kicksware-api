$(document).ready(function () {
	"use strict";
});

function addSearchResultCell(item) {
	let searchCell = document.createElement("DIV");
	searchCell.className = "search-cell";

	let thumb = document.createElement("DIV");
	thumb.className = "thumb";
	let image = document.createElement("IMG");
	image.setAttribute("src", item.imageLink);
	thumb.append(image);

	let info = document.createElement("DIV");
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
	info.append(brandTitle, modelTitle, skuCode);

	searchCell.append(thumb, info);
	$(".search-grid").append(searchCell);
}

function initAutoSearch(actionUrl) {
	$("#auto-search").autocomplete({
		source: function (request) {
			$.ajax({
				url: `${actionUrl}?prefix=${request.term}`,
				type: "GET",
				contentType: "application/json",
				dataType: "json",
				success: function (data) {
					for (let item of data) {
						addSearchResultCell(item);
					}
				}
			});
		}
	});
}