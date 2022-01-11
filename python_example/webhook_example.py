from flask import Flask, request, json

app = Flask(__name__)


@app.route('/')
def hello():
    return 'Nyahello'


@app.route('/github_issue', methods=['POST'])
def github_issue():
    data = request.json
    print(f"Issue {data['issue']['title']} {data['action']}")
    print(data['issue']['body'])
    print(data['issue']['url'])
    return data


if __name__ == '__main__':
    app.run(debug=True)
