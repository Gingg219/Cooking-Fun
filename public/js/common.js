function showToastSuccess(data) {
    $('#titleSuccess').text(data.title);
    $('#contentSuccess').text(data.message);
    $('#toastSuccess').toast('show');
}
    
function showToastFail(data) {
    $('#titleFail').text(data.title);
    $('#contentFail').text(data.message);
    $('#toastFail').toast('show');
}

$('.custom_select2').select2({
    dropdownParent: $('.modal'),
    width: "100%",
    selectionCssClass: ":all:",
    allowClear: true,
    tags: true
});
function scrollTopPaginate() {
    $('html, main').animate({ scrollTop: 0 }, 800);
}

function dateTimeFormat(dateTime) {
    var dateTimeObj = {
        "year": dateTime.getFullYear(),
        "month": dateTime.getMonth() + 1,
        "day": dateTime.getDate(),
    };
    for (var k in dateTimeObj) {
        if (dateTimeObj[k] < 10) {
            dateTimeObj[k] = "0" + dateTimeObj[k];
        }
    }
    var dateTimeString = dateTimeObj.day + "-" + dateTimeObj.month + "-" + dateTimeObj.year;
    return {"obj": dateTimeObj, "string": dateTimeString};
}

function formatDate($id, $startDate, $maxDate, $minDate) {
    $id.prop('readonly', true);
    $id.daterangepicker({
        showDropdowns: true,
        singleDatePicker: true,
        locale: {
            format: "DD-MM-YYYY",
            daysOfWeek: ["CN", "Thứ 2", "Thứ 3", "Thứ 4", "Thứ 5", "Thứ 6", "Thứ 7"],
            monthNames: ["Tháng 01", "Tháng 02", "Tháng 03", "Tháng 04", "Tháng 05", "Tháng 06", "Tháng 07", "Tháng 08", "Tháng 09", "Tháng 10", "Tháng 11", "Tháng 12"],
            firstDay: 1
        },
        minDate: $minDate,
        maxDate: $maxDate,
    }).on('apply.daterangepicker', function(ev, picker) {
        picker.element.val(dateTimeFormat(picker.startDate._d).string);
    });
    $id.data("daterangepicker").setStartDate($startDate);
}