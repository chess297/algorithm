class Array<T> {
  late List<T> _data;
  Array({required List<T> data}) {
    this._data = data;
  }

  add(T e) {
    this._data.add(e);
  }

  T getLast() {
    return this._data.last;
  }

  T? getFirst() {
    return this._data.firstOrNull;
  }
}
