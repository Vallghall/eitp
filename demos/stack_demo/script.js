document.addEventListener("DOMContentLoaded", () => {
    const stack = document.getElementById("stack");
    const pushButton = document.getElementById("pushButton");
    const popButton = document.getElementById("popButton");
    const err = document.getElementById("error")

    let elementCount = 0;
    const names = ["main", "foo", "bar", "baz", "qux"]

    pushButton.addEventListener("click", () => {
        if (elementCount >= 5) {
            err.style.color = "#000000";
            return;
        }

        const element = document.createElement("div");
        element.className = "element";
        element.innerText = names[elementCount];
        stack.appendChild(element);
        elementCount++;
    });

    popButton.addEventListener("click", () => {
        err.style.color = "#FFFFFF";
        const elements = stack.getElementsByClassName("element");
        if (elements.length > 0) {
            const lastElement = elements[elements.length - 1];
            lastElement.style.animation = "fadeOut 0.5s ease-in-out";
            setTimeout(() => {
                stack.removeChild(lastElement);
                elementCount--;
            }, 490);
        }
    });
});
