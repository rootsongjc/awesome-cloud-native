(function () {
    "use strict";

    var GITHUB_REPO_RE = /^https:\/\/github\.com\/([^/]+\/[^/)]+?)(?:[/?].*)?$/;

    function loadStatus(callback) {
        var xhr = new XMLHttpRequest();
        xhr.open("GET", "./status.json", true);
        xhr.responseType = "json";
        xhr.onload = function () {
            if (xhr.status === 200 && xhr.response) {
                callback(xhr.response);
            }
        };
        xhr.send();
    }

    function createDot(status) {
        var dot = document.createElement("span");
        dot.className = "status-dot status-dot--" + status;
        dot.title = status.charAt(0).toUpperCase() + status.slice(1);
        return dot;
    }

    function createText(text) {
        return document.createTextNode(text);
    }

    function injectDots(data) {
        var projects = data.projects;
        var links = document.querySelectorAll("#content a[href]");
        var matched = 0;

        links.forEach(function (link) {
            var match = link.href.match(GITHUB_REPO_RE);
            if (!match) return;

            // Only inject into list items (project entries), skip headings/footer/etc
            var li = link.closest("li");
            if (!li) return;

            // Skip Table of Contents
            var parentUl = li.parentElement;
            if (parentUl && parentUl.previousElementSibling && parentUl.previousElementSibling.textContent === "Table of Contents") return;

            var repo = match[1];
            var info = projects[repo];
            if (!info) return;

            var status = info.status;
            var dot = createDot(status);
            li.insertBefore(dot, li.firstChild);

            if (status !== "active") {
                link.classList.add("status-link--" + status);
            }

            matched++;
        });

        return matched;
    }

    function createStat(status, count, label) {
        var stat = document.createElement("div");
        stat.className = "status-summary__stat";

        var countEl = document.createElement("span");
        countEl.className = "status-summary__count status-summary__count--" + status;
        countEl.appendChild(createText(count));

        var labelEl = document.createElement("span");
        labelEl.className = "status-summary__label";
        labelEl.appendChild(createText(label));

        stat.appendChild(countEl);
        stat.appendChild(labelEl);
        return stat;
    }

    function injectSummary(data) {
        var content = document.getElementById("content");
        if (!content) return;

        var summary = data.summary;
        var updated = data.updated_at ? data.updated_at.split("T")[0] : "unknown";

        var banner = document.createElement("div");
        banner.className = "status-summary";

        var title = document.createElement("p");
        title.className = "status-summary__title";
        title.appendChild(createText("Project Status Overview · Updated " + updated));

        var stats = document.createElement("div");
        stats.className = "status-summary__stats";

        stats.appendChild(createStat("active", summary.active || 0, "Active"));
        stats.appendChild(createStat("archived", summary.archived || 0, "Archived"));
        stats.appendChild(createStat("deleted", summary.deleted || 0, "Deleted / 404"));
        stats.appendChild(createStat("inactive", summary.inactive || 0, "Inactive (>2y)"));

        banner.appendChild(title);
        banner.appendChild(stats);
        content.insertBefore(banner, content.firstChild);
    }

    function createLegendItem(status, text) {
        var item = document.createElement("span");
        item.className = "status-legend__item";
        item.appendChild(createDot(status));
        item.appendChild(createText(" " + text));
        return item;
    }

    function injectLegend() {
        var content = document.getElementById("content");
        if (!content) return;

        var legend = document.createElement("div");
        legend.className = "status-legend";

        legend.appendChild(createLegendItem("active", "Active"));
        legend.appendChild(createLegendItem("archived", "Archived"));
        legend.appendChild(createLegendItem("deleted", "Deleted / 404"));
        legend.appendChild(createLegendItem("inactive", ">2 years inactive"));

        content.appendChild(legend);
    }

    // Entry point
    loadStatus(function (data) {
        injectSummary(data);
        var count = injectDots(data);
        injectLegend();
        console.log("[status-checker] Injected " + count + " status dots");
    });
})();
