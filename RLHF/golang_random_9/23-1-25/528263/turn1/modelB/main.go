func mapSliceUsingReflection<T, U>(_ slice: [T], _ transform: @escaping (T) -> U) -> [U] {
    var resultSlice: [U] = []

    let sliceType = type(of: slice)
    let elementType = sliceType.Element.self

    let resultSliceType = Array<U>.self
    let resultElementType = resultSliceType.Element.self

    let sliceMirror = Mirror(reflecting: slice)

    for child in sliceMirror.children {
        if let element = child.value as? T {
            let transformedElement = transform(element)
            resultSlice.append(transformedElement)
        }
    }

    return resultSlice
}