"""INTENTIONALLY FLAWED agent loop — candidates must replace its trust model."""


def answer(model, toxindex, question):
    tools = toxindex.list_tools()
    messages = [{"role": "user", "content": question}]
    while True:  # Fault: no round, time, token, cost or cancellation budget.
        response = model.generate(messages, tools=tools)
        if not response.tool_calls:
            # Fault: no structured claims or mechanical citation validation.
            return {"answer": response.text, "citations": response.citations}
        for call in response.tool_calls:
            result = toxindex.call(call.name, call.arguments)
            # Fault: retrieved instructions are handed to the model as trusted.
            messages.append({"role": "user", "content": result})

