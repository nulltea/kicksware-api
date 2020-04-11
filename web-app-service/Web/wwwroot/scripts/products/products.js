$(document).ready(function() {
    function initIsotope() {
        var sortingButtons = $('.product_sorting_btn');
        var sortNums = $('.num_sorting_btn');

        if ($('.product_grid').length) {
            var grid = $('.product_grid').isotope({
                itemSelector: '.product',
                layoutMode: 'fitRows',
                fitRows:
                {
                    gutter: 30
                },
                getSortData:
                {
                    price: function(itemElement) {
                        var priceEle = $(itemElement).find('.product_price').text().replace('$', '');
                        return parseFloat(priceEle);
                    },
                    name: '.product_name',
                    stars: function(itemElement) {
                        var starsEle = $(itemElement).find('.rating');
                        var stars = starsEle.attr("data-rating");
                        return stars;
                    }
                },
                animationOptions:
                {
                    duration: 750,
                    easing: 'linear',
                    queue: false
                }
            });

            // Sort based on the value from the sorting_type dropdown
            sortingButtons.each(function() {
                $(this).on('click',
                    function() {
                        var parent = $(this).parent().parent().find('.sorting_text');
                        parent.text($(this).text());
                        var option = $(this).attr('data-isotope-option');
                        option = JSON.parse(option);
                        grid.isotope(option);
                    });
            });
        }
    }

});