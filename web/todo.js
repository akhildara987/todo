window.onload = function () {



    $(document).on('click', '#addtask', function (e) {
        //handler code here
        var jsonData = {};
        jsonData['title'] = $("#title").val().trim();
        jsonData['description'] = $("#description").val().trim();
        jsonData['todoPriority'] = $("#pri").val().trim();
        jsonData['enddate'] = $("#endtime").val().trim();
        var id = sessionStorage.getItem('id');

        console.log(JSON.stringify(jsonData));

        $.ajax({
            url: '/api/todo?id=' + id,
            contentType: 'application/json',
            type: 'post',
            data: JSON.stringify(jsonData),

            success: function (response) {
                console.log(response);
                // console.log(response.redirecturl);s
                var msg = "";
                alert("todo added");
                location.reload();

            }
        });

    });


    $(document).on('click', '#e_addtask', function (e) {
        //handler code here
        var jsonData = {};
        jsonData['title'] = $("#e_title").val().trim();
        jsonData['description'] = $("#e_description").val().trim();
        jsonData['todoPriority'] = $("#e_pri").val().trim();
        jsonData['enddate'] = $("#e_endtime").val().trim();
        jsonData['todoID'] = sessionStorage.getItem('ecid');

        var id = sessionStorage.getItem('id');

        console.log(JSON.stringify(jsonData));

        $.ajax({
            url: '/api/todo?id=' + id,
            contentType: 'application/json',
            type: 'put',
            data: JSON.stringify(jsonData),

            success: function (response) {
                console.log(response);
                // console.log(response.redirecturl);s
                var msg = "";
                alert("todo updated");
                location.reload();

            }
        });

    });


    $(document).on('click', '#editthis', function (e) {
        var todoID = $(this).attr('etodoid');
        sessionStorage.setItem("ecid", todoID);
        $('#myModal').modal('show');


    });


    $(document).on('click', '#tododelete', function (e) {
        //handler code here
        var id = sessionStorage.getItem('id');
        var todoID = $(this).attr('todoid');

        $.ajax({
            url: '/api/todo?id=' + id + "&todoid=" + todoID,
            contentType: 'application/json',
            type: 'delete',
            success: function (response) {
                console.log(response);
                // console.log(response.redirecturl);s
                var msg = "";
                alert("todo deleted");
                location.reload();

            }
        });

    });

    $(document).on('click', '#tocompleted', function (e) {
        //handler code here
        var id = sessionStorage.getItem('id');
        var todoID = $(this).attr('todoid');

        $.ajax({
            url: '/api/todo/completed?id=' + id + "&todoid=" + todoID,
            contentType: 'application/json',
            type: 'put',
            success: function (response) {
                console.log(response);
                // console.log(response.redirecturl);s
                var msg = "";
                alert("todo completed");
                location.reload();

            }
        });

    });



    var id = sessionStorage.getItem("id");
    console.log(id)
    $.ajax({
        url: '/api/todo?id=' + id,
        contentType: 'application/json',
        type: 'get',
        success: function (response) {
            console.log(response);
            var msg = "";
            var data;

            $.each(response, function (key, value) {
                console.log(value.title);
                var iscomp = ``;
                var icolor='text-warning';
                var bcolor='border-warning';

                if (value.is_completed === 'true') {
                    iscomp += `<i class="fa fa-square-o text-primary btn m-0 p-0 d-none" data-toggle="tooltip" data-placement="bottom" title="Mark as complete"></i>
                    <i class="fa fa-check-square-o text-primary btn m-0 p-0" data-toggle="tooltip" data-placement="bottom" title="Mark as todo"></i>`;
                     icolor='text-warning-green';
                     bcolor='border-warning-green';
                } else {
                    iscomp += `<i class="fa fa-square-o text-primary btn m-0 p-0" data-toggle="tooltip" data-placement="bottom" title="Mark as complete"></i>`;
                }
                data = `<div class="row px-3 align-items-center todo-item rounded">
            <div class="col-auto m-1 p-0 d-flex align-items-center">
                <h2 class="m-0 p-0" id="tocompleted" todoid="${value.todoID}">
                ${iscomp}
                    </h2>
            </div>
            <div class="col px-1 m-1 d-flex align-items-center">
                <input type="text" class="form-control form-control-lg border-0 edit-todo-input bg-transparent rounded px-3" readonly value="${value.title}" title="${value.title}" />
                <input type="text" class="form-control form-control-lg border-0 edit-todo-input rounded px-3 d-none" value="${value.title}" />
            </div>
            <div class="col-auto m-1 p-0 px-3">
            <div class="row">
                <div class="col-auto d-flex align-items-center rounded bg-white border ${bcolor}">
                    <i class="fa fa-hourglass-2 my-2 px-2 ${icolor} btn" data-toggle="tooltip" data-placement="bottom" title="" data-original-title="Due on date"></i>
                    <h6 class="text my-2 pr-2">${value.enddate}</h6>

                </div>
            </div>
        </div>
          
            <div class="col-auto m-1 p-0 todo-actions">
                <div class="row d-flex align-items-center justify-content-end">
                <h5 class="m-0 p-0 px-2">
                ${value.TodoPriority}
                    </h5>
                    <h5 class="m-0 p-0 px-2" id="editthis" etodoid="${value.todoID}">
                        <i class="fa fa-pencil text-info btn m-0 p-0" data-toggle="tooltip" data-placement="bottom" title="Edit todo"></i>
                    </h5>
                    <h5 class="m-0 p-0 px-2" id="tododelete" todoid="${value.todoID}">
                        <i class="fa fa-trash-o text-danger btn m-0 p-0"  data-toggle="tooltip" data-placement="bottom" title="Delete todo"></i>
                    </h5>
                </div>
                <div class="row todo-created-info">
                    <div class="col-auto d-flex align-items-center pr-2">
                        <i class="fa fa-info-circle my-2 px-2 text-black-50 btn" data-toggle="tooltip" data-placement="bottom" title="" data-original-title="Created date"></i>
                        <label class="date-label my-2 text-black-50">${(value.enddate)}</label>
                    </div>
                </div>
            </div>
        </div>`
                $(data).appendTo('#todolist');

            });

        }
    });





    bootlint.showLintReportForCurrentDocument([], {
        hasProblems: false,
        problemFree: false
    });


    $('[data-toggle="tooltip"]').tooltip();
    function formatDate(date) {
        return (
            date.getDate() +
            "/" +
            (date.getMonth() + 1) +
            "/" +
            date.getFullYear()
        );
    }


    var currentDate = formatDate(new Date());

    $(".due-date-button").datepicker({
        format: "dd/mm/yyyy",
        autoclose: true,
        todayHighlight: true,
        startDate: currentDate,
        orientation: "bottom right"
    });

    $(".due-date-button").on("click", function (event) {
        $(".due-date-button")
            .datepicker("show")
            .on("changeDate", function (dateChangeEvent) {
                $(".due-date-button").datepicker("hide");
                $(".due-date-label").text(formatDate(dateChangeEvent.date));
            });
    });
};
