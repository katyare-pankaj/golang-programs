import Foundation
// Define a complex data type struct
struct Point3D {
    let x: Double
    let y: Double
    let z: Double
}

// Traditional fmt.Sprintf usage with Point3D
func traditionalFormat(_ p: Point3D) -> String {
    return fmt.Sprintf("Point3D(%.2f, %.2f, %.2f)", p.x, p.y, p.z)
}

// Enhanced version of fmt.Sprintf using fmt.State interface
func enhancedFormat(_ p: Point3D, f fmt.State, c rune) {
    fmt.Fprintf(f, "Point3D(%.2f, %.2f, %.2f)", p.x, p.y, p.z)
}

func main() {
    p := Point3D{x: 1.2345, y: 6.7890, z: 12.3456}

    // Calculate the baseline time for traditional formatting
    start := time.Now()
    for i := 0; i < 1_000_000; i++ {
        _ = traditionalFormat(p)
    }
    baselineTime := time.Since(start)

    // Calculate the enhanced time for enhanced formatting
    start = time.Now()
    for i := 0; i < 1_000_000; i++ {
        _ = fmt.Sprintf("%s", enhancedFormat(p))
    }
    enhancedTime := time.Since(start)

    fmt.Printf("Traditional Format Time: %s\n", baselineTime)
    fmt.Printf("Enhanced Format Time: %s\n", enhancedTime)
}