from reportlab.lib.pagesizes import letter
from reportlab.platypus import SimpleDocTemplate, Paragraph, HRFlowable
from reportlab.lib.styles import getSampleStyleSheet, ParagraphStyle
from reportlab.lib import colors
from reportlab.lib.units import inch
from pathlib import Path

# Save next to this script
pdf_path = Path(__file__).parent / "Amaresh_Thakur_Amazon_SDE_Intern_Resume.pdf"

name = "MUKESH AMARESH THAKUR"
headline = "Software Development Engineer Intern Candidate | Go (Golang)"
contact = "mukeshamareshthakur@gmail.com | Nepal"
github = "github.com/amareshthakur"
linkedin = "linkedin.com/in/mukesh-amaresh-thakur-327a45230"
project_url = "github.com/thaakur-u/attendance_system"

summary = ("Electronics, Communication and Information Engineering undergraduate with hands-on academic project "
           "experience building a Go-based attendance management system and an embedded-camera emotion detection "
           "prototype. Interested in backend engineering, data structures and algorithms, databases, and the "
           "integration of software with hardware.")

skills = [
    ("Languages", "Go (Golang), Python"),
    ("Backend & Security", "Gin, REST API development, PASETO (project implementation)"),
    ("Databases", "PostgreSQL"),
    ("Frontend", "Next.js, Tailwind CSS, shadcn/ui"),
    ("Embedded / IoT", "ESP32, ESP32-CAM, fingerprint scanner, RFID reader"),
    ("ML / Computer Vision", "CNN, ResNet, emotion classification"),
    ("Core CS", "Data structures and algorithms, problem-solving; add only topics you can confidently discuss"),
]

projects = [
    ("Automated Class Attendance System | Major Academic Project",
     "Go, Gin, PostgreSQL, PASETO, Next.js, Tailwind CSS, shadcn/ui, ESP32, fingerprint scanner, RFID reader",
     [
         "Developed an attendance-system project designed to record attendance automatically according to each class time slot.",
         "Built the backend using Go and Gin with PostgreSQL for data storage and PASETO for token-based authentication.",
         "Built the web interface with Next.js, Tailwind CSS, and shadcn/ui; integrated fingerprint and RFID hardware with an ESP32-based workflow."
     ]),
    ("Facial Emotion Detection | Minor Academic Project",
     "Python, CNN, ResNet, ESP32-CAM, Kaggle emotion-detection dataset",
     [
         "Developed a prototype that captures a face image using an ESP32-CAM and classifies the facial emotion with a trained model.",
         "Used a CNN/ResNet-based approach trained on a Kaggle emotion-detection dataset."
     ])
]

pdf_styles = getSampleStyleSheet()
pdf_styles.add(ParagraphStyle(name="ResumeName", parent=pdf_styles["Normal"], fontName="Helvetica-Bold",
                              fontSize=15, leading=16, alignment=1, textColor=colors.HexColor("#1F374B"), spaceAfter=1))
pdf_styles.add(ParagraphStyle(name="ResumeHeadline", parent=pdf_styles["Normal"], fontName="Helvetica-Bold",
                              fontSize=8.8, leading=10, alignment=1, spaceAfter=1))
pdf_styles.add(ParagraphStyle(name="ResumeContact", parent=pdf_styles["Normal"], fontName="Helvetica",
                              fontSize=7.4, leading=8.5, alignment=1, spaceAfter=4))
pdf_styles.add(ParagraphStyle(name="SectionHeading", parent=pdf_styles["Normal"], fontName="Helvetica-Bold",
                              fontSize=9, leading=10, textColor=colors.HexColor("#1F374B"), spaceBefore=4, spaceAfter=1))
pdf_styles.add(ParagraphStyle(name="ResumeBody", parent=pdf_styles["Normal"], fontName="Helvetica",
                              fontSize=7.6, leading=9, spaceAfter=1))
pdf_styles.add(ParagraphStyle(name="ResumeBullet", parent=pdf_styles["Normal"], fontName="Helvetica",
                              fontSize=7.4, leading=8.6, leftIndent=10, firstLineIndent=-7, spaceAfter=0.7))
pdf_styles.add(ParagraphStyle(name="ProjectTitle", parent=pdf_styles["Normal"], fontName="Helvetica-Bold",
                              fontSize=8, leading=9, spaceBefore=2, spaceAfter=0.5))

story = [
    Paragraph(name, pdf_styles["ResumeName"]),
    Paragraph(headline, pdf_styles["ResumeHeadline"]),
    Paragraph(f"{contact} | {github} | {linkedin}", pdf_styles["ResumeContact"]),
]

def section(title):
    story.append(Paragraph(title.upper(), pdf_styles["SectionHeading"]))
    story.append(HRFlowable(width="100%", thickness=0.45, color=colors.HexColor("#AAB5BF"), spaceAfter=2))

def bullet(text):
    story.append(Paragraph("• " + text, pdf_styles["ResumeBullet"]))

section("Professional Summary")
story.append(Paragraph(summary, pdf_styles["ResumeBody"]))

section("Technical Skills")
for label, value in skills:
    story.append(Paragraph(f"<b>{label}:</b> {value}", pdf_styles["ResumeBody"]))

section("Projects")
for title, tech, bullets in projects:
    story.append(Paragraph(title, pdf_styles["ProjectTitle"]))
    story.append(Paragraph(f"<b>Technologies:</b> {tech}", pdf_styles["ResumeBody"]))
    for b in bullets:
        bullet(b)
story.append(Paragraph(f"<b>Project repository:</b> {project_url}", pdf_styles["ResumeBody"]))

section("Education")
story.append(Paragraph("<b>Bachelor of Engineering — Electronics, Communication and Information Engineering</b>", pdf_styles["ResumeBody"]))
story.append(Paragraph("IOE, Paschimanchal Campus, Tribhuvan University, Nepal | Expected graduation: January 2027 (confirm) | Academic score: 65%", pdf_styles["ResumeBody"]))

section("Additional Information")
story.append(Paragraph("Experience: Academic/self-directed projects; no formal internship reported.", pdf_styles["ResumeBody"]))

pdf = SimpleDocTemplate(str(pdf_path), pagesize=letter, rightMargin=0.55*inch, leftMargin=0.55*inch,
                        topMargin=0.38*inch, bottomMargin=0.38*inch)
pdf.build(story)

print("Created:", pdf_path, pdf_path.stat().st_size, "bytes")