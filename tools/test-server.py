from flask import Flask, request

APP = Flask(__name__)

@APP.route('/')
def main():
    return 'Main page'

@APP.route('/post', methods=['POST'])
def post():
    d = {}
    for k, v in request.form.items():
        d[k] = v
    results = '{}\n'.format(d)
    print(results)
    return results

if __name__ == '__main__':
    APP.run(
        host='0.0.0.0',
        port=5000,
        debug=True,
    )
