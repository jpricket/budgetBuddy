function handleEnter() {
    if (event.keyCode === 13) {
        event.preventDefault();
        return false;
    }
}

function saveData() {
    $.ajax({
        url: "/budgets",
        type: "POST",
        data: JSON.stringify(getTableData()),
        contentType: "application/json; charset=utf-8",
        success: function (result) {
            alert("Data Saved");
            $('#testDiv').text(result);
        },
        error: function (xhr, ajaxOptions, thrownError) {
            alert(xhr.status);
            alert(thrownError);
        }
    });
}

function getTableData() {
    const month = "{{.Month}}";
    const year = "{{.Year}}";
    const budgetItems = [];

    const table = $("#items");
    table.find('tr.data').each(function (i, el) {
        const $tds = $(this).find('td');
        const name = $tds.eq(0).text();
        const budgetAmount = $tds.eq(1).text();
        const actualAmount = $tds.eq(2).text();
        budgetItems.push({
            name: name,
            budgetAmount: parseFloat(budgetAmount),
            actualAmount: parseFloat(actualAmount)
        })
    });

    $('#testDiv').text(JSON.stringify(budgetItems));

    return {
        month: month,
        year: parseInt(year),
        budgetItems: budgetItems
    };
}
